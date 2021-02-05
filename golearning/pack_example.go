package main

import (
	"fmt"

	"go.etcd.io/etcd/clientv3/namespace"
	"k8s.io/apimachinery/pkg/util/uuid"
)

type Project struct {
	Id uuid
	Name string
	NameSpace namespace
}
func main() {
	projects := make(map[string]Project)
	var p Project
	p.Name := "myapp"
	p.Id := NewUUID()
	projects = {"p.Name": {p.Id, p.Name, nil},}
	fmt.Println (projects["myapp"])

}
