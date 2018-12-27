package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"flag"
	"path/filepath"
	"log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AutoGenerated struct {
	Status string `json:"status"`
	Data   struct {
		Alerts []struct {
			Labels struct {
				Alertname           string `json:"alertname"`
				App                 string `json:"app"`
				Container           string `json:"container"`
				Instance            string `json:"instance"`
				Job                 string `json:"job"`
				K8SApp              string `json:"k8s_app"`
				KubernetesName      string `json:"kubernetes_name"`
				KubernetesNamespace string `json:"kubernetes_namespace"`
				Namespace           string `json:"namespace"`
				Pod                 string `json:"pod"`
				Severity            string `json:"severity"`
				Deployment          string `json:"deployment"`
			} `json:"labels"`
			Annotations struct {
				Message string `json:"message"`
			} `json:"annotations"`
			State    string    `json:"state"`
			ActiveAt time.Time `json:"activeAt"`
			Value    float64   `json:"value"`
		} `json:"alerts"`
	} `json:"data"`
}

func IndexNum(weburl string) (jsonstr string) {
	// Request Prometheus API
	resp, err := http.Get(weburl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	Body, err := ioutil.ReadAll(resp.Body)
	jsonStr := string(Body)
	return jsonStr
}


func Deployment(jsonstr string,clientset *kubernetes.Clientset) {
	var autogenerated AutoGenerated
	var mapList map[string]string

	err := json.Unmarshal([]byte(jsonstr), &autogenerated)
	if err != nil {
		panic(err)
	}
	log.Printf("autogenerated: %T\n",autogenerated)
	// Total number of alarm entries
	log.Printf("total entry:%d\n",len(autogenerated.Data.Alerts))
	mapList = make(map[string]string)

	for i := 0; i < len(autogenerated.Data.Alerts); i++ {
		if autogenerated.Data.Alerts[i].Labels.Alertname == "KubeDeploymentReplicasMismatch" {
			log.Printf("namespace null : %v,%v\n",autogenerated.Data.Alerts[i].Labels.Namespace,autogenerated.Data.Alerts[i].Labels.Deployment) 
			mapList[autogenerated.Data.Alerts[i].Labels.Namespace] = autogenerated.Data.Alerts[i].Labels.Deployment
		}

	}

	if len(mapList) > 0 {
		log.Println("Need to deal with the deployment:", mapList,len(mapList))
		CleanDeployment(mapList,clientset)
		CleanIngress(mapList,clientset)
		CleanService(mapList,clientset)
	} else {
		log.Println("There is no need to deal with the Deployment.")
	}
}

func CleanDeployment(mapList map[string]string,clientset *kubernetes.Clientset) {
	for namespace,deploymentname := range mapList {
		fmt.Printf("delete deployment:[%v %v]\n",namespace,deploymentname)

		deploymentsClient := clientset.AppsV1().Deployments(namespace)
		deletePolicy := metav1.DeletePropagationForeground
		if err := deploymentsClient.Delete(deploymentname, &metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
		time.Sleep(time.Duration(1)*time.Second)
	}
}

func CleanIngress(mapList map[string]string,clientset *kubernetes.Clientset) {
	for namespace,deploymentname := range mapList {  
		// GET
		// ingress, err := clientset.ExtensionsV1beta1().Ingresses(namespace).Get(deploymentname, metav1.GetOptions{})
		// DELETE 
		log.Printf("delete ingress:[%v %v]\n",namespace,deploymentname)
	    err := clientset.ExtensionsV1beta1().Ingresses(namespace).Delete(deploymentname, &metav1.DeleteOptions{})
		if err != nil {
			panic(err)
		}
	}
}

func CleanService(mapList map[string]string,clientset *kubernetes.Clientset) {
	for namespace,deploymentname := range mapList {  
		// GET
		// service, err := clientset.CoreV1().Services(namespace).Get(deploymentname, metav1.GetOptions{})
		// DELETE
		log.Printf("delete services:[%v %v]\n",namespace,deploymentname)
		err := clientset.CoreV1().Services(namespace).Delete(deploymentname, &metav1.DeleteOptions{})
		if err != nil {
			panic(err)
		}
	}	
}

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	var webURL = flag.String("web_url", "", "Get http url (default return json)")
	flag.Parse()

	if *webURL == "" {
		flag.Usage()
	}

	
	jsonstr := IndexNum(*webURL)

	fmt.Println(jsonstr)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	Deployment(jsonstr,clientset)
}