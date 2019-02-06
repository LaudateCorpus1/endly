package v1



import (
	"fmt"
"errors"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	vvc "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"	
)

/*autogenerated contract adapter*/

//EndpointsCreateRequest represents request
type EndpointsCreateRequest struct {
  service_ v1.EndpointsInterface
   *vvc.Endpoints
}

//EndpointsUpdateRequest represents request
type EndpointsUpdateRequest struct {
  service_ v1.EndpointsInterface
   *vvc.Endpoints
}

//EndpointsDeleteRequest represents request
type EndpointsDeleteRequest struct {
  service_ v1.EndpointsInterface
  Name string
   *metav1.DeleteOptions
}

//EndpointsDeleteCollectionRequest represents request
type EndpointsDeleteCollectionRequest struct {
  service_ v1.EndpointsInterface
   *metav1.DeleteOptions
  ListOptions metav1.ListOptions
}

//EndpointsGetRequest represents request
type EndpointsGetRequest struct {
  service_ v1.EndpointsInterface
  Name string
   metav1.GetOptions
}

//EndpointsListRequest represents request
type EndpointsListRequest struct {
  service_ v1.EndpointsInterface
   metav1.ListOptions
}

//EndpointsWatchRequest represents request
type EndpointsWatchRequest struct {
  service_ v1.EndpointsInterface
   metav1.ListOptions
}

//EndpointsPatchRequest represents request
type EndpointsPatchRequest struct {
  service_ v1.EndpointsInterface
  Name string
  Pt types.PatchType
  Data []byte
  Subresources []string
}


func init() {
	register(&EndpointsCreateRequest{})
	register(&EndpointsUpdateRequest{})
	register(&EndpointsDeleteRequest{})
	register(&EndpointsDeleteCollectionRequest{})
	register(&EndpointsGetRequest{})
	register(&EndpointsListRequest{})
	register(&EndpointsWatchRequest{})
	register(&EndpointsPatchRequest{})
}


func (r * EndpointsCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.EndpointsInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.EndpointsInterface", service)
	}
	return nil
}

func (r * EndpointsCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.Endpoints)
	return result, err	
}

func (r * EndpointsCreateRequest) GetId() string {
	return "v1.Endpoints.Create";	
}

func (r * EndpointsUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.EndpointsInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.EndpointsInterface", service)
	}
	return nil
}

func (r * EndpointsUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.Endpoints)
	return result, err	
}

func (r * EndpointsUpdateRequest) GetId() string {
	return "v1.Endpoints.Update";	
}

func (r * EndpointsDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.EndpointsInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.EndpointsInterface", service)
	}
	return nil
}

func (r * EndpointsDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name,r.DeleteOptions)
	return result, err	
}

func (r * EndpointsDeleteRequest) GetId() string {
	return "v1.Endpoints.Delete";	
}

func (r * EndpointsDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.EndpointsInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.EndpointsInterface", service)
	}
	return nil
}

func (r * EndpointsDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions,r.ListOptions)
	return result, err	
}

func (r * EndpointsDeleteCollectionRequest) GetId() string {
	return "v1.Endpoints.DeleteCollection";	
}

func (r * EndpointsGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.EndpointsInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.EndpointsInterface", service)
	}
	return nil
}

func (r * EndpointsGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name,r.GetOptions)
	return result, err	
}

func (r * EndpointsGetRequest) GetId() string {
	return "v1.Endpoints.Get";	
}

func (r * EndpointsListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.EndpointsInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.EndpointsInterface", service)
	}
	return nil
}

func (r * EndpointsListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err	
}

func (r * EndpointsListRequest) GetId() string {
	return "v1.Endpoints.List";	
}

func (r * EndpointsWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.EndpointsInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.EndpointsInterface", service)
	}
	return nil
}

func (r * EndpointsWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err	
}

func (r * EndpointsWatchRequest) GetId() string {
	return "v1.Endpoints.Watch";	
}

func (r * EndpointsPatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.EndpointsInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.EndpointsInterface", service)
	}
	return nil
}

func (r * EndpointsPatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name,r.Pt,r.Data,r.Subresources...)
	return result, err	
}

func (r * EndpointsPatchRequest) GetId() string {
	return "v1.Endpoints.Patch";	
}
