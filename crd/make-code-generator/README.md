# Make Code Generator

- CustomResources in golang have to implement the ```runtime.Object interface```.
- ```runtime.Object``` types must have DeepCopy methods.
- deepcopy-gen: creates a method ```func (t* T) DeepCopy() *T``` for each type T.
- client-gen: creates typed clientsets for CustomResources APIGroups.
- informer-gen: creates informers for CustomResources which offer an event based interface to react on
  changes of CustomResources on the server.
- lister-gen: creates listers for CustomResources which offer a read-only caching layer for GET and 
  LIST requests.
- There are two kind of tags:
  - Global tags above ```package``` in ```doc.go```.
  - Local tags above a type that is processed.
- Tags in general have the shape ```// +tag-name``` or ```// +tag-name=value```, that is, they are written into comments.
- Global tags are written into the ```doc.go``` file of a package. A typical ```pkg/apis/<apigroup>/<version>/doc.go```
  looks like this:
  ```
    // +k8s:deepcopy-gen=package,register
    // Package v1 is the v1 version of the API.
    // +groupName=example.com
    package v1
  ```
- Local tags are written either directly above an API type or in the second comment block above it. Example of types.go:
  ```
  // +genclient
  // +genclient:noStatus
  // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
  // Database describes a database.
  type Database struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    
    Spec DatabaseSpec `json:"spec"`
    }
    
    
    // DatabaseSpec is the spec for a Foo resource
    type DatabaseSpec struct {
    User string `json:"user"`
    Password string `json:"password"`
    Encoding string `json:"encoding,omitempty"`
    }
    
    
    // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
    
    
    // DatabaseList is a list of Database resources
    type DatabaseList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata"`
    
    
    Items []Database `json:"items"`
    }
  ```
- 