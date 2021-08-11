package controllers

import (
	"DemoProject/api/auth"
	"DemoProject/api/models"
	"DemoProject/api/responses"
	"DemoProject/api/utils"
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"io/ioutil"
	"net/http"
)

func (s *Server) CreateResource(w http.ResponseWriter,  r *http.Request) {

	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	//fmt.Println("String", str)
	var resources []models.Resource
	resources = fetchResourceFromAws()
	for i, _ := range resources {
		resourceCreated, err := resources[i].SaveResource(s.DB)
		if err != nil {
			formattedError := utils.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}
		//w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, resourceCreated.InstanceId))
		responses.JSON(w, http.StatusCreated, resourceCreated)
	}

}

func fetchResourceFromAws() []models.Resource{
	var resources []models.Resource
    var resource = models.Resource{}
	sess := ConnectAws("","")
	ec2Svc := ec2.New(sess)
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		//fmt.Println("Success", result)
		for idx, _ := range result.Reservations {
			for _, inst := range result.Reservations[idx].Instances {
				resource.ImageId = *inst.ImageId
				resource.InstanceId = *inst.InstanceId
				resource.InstanceType = *inst.InstanceType
				resource.RootDeviceType = *inst.RootDeviceType
				resource.RootDeviceName = *inst.RootDeviceName
				resource.PrivateDnsName = *inst.PrivateDnsName
			}
			resources = append(resources, resource)
	    }

	}
	return resources
}

func (s *Server) GetResources(w http.ResponseWriter, r *http.Request) {

	userId, err :=  auth.ExtractTokenID(r)
	fmt.Printf("Userid ExtractTokenID method:%v", userId)
	resource := models.Resource{}
	resources, err := resource.FindAllResources(s.DB,userId)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, resources))
	responses.JSON(w, http.StatusOK, resources)
}
