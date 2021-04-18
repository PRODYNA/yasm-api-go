# \ProjectApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonProject**](ProjectApi.md#AddPersonProject) | **Post** /persons/{personId}/projects/{projectId} | Add Project to a Person
[**ConfirmSkill**](ProjectApi.md#ConfirmSkill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /organizations/{organizationId}/projects | Create a Project in an Organization
[**DeleteConfirmation**](ProjectApi.md#DeleteConfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
[**DeletePersonProject**](ProjectApi.md#DeletePersonProject) | **Delete** /persons/{personId}/projects/{projectId} | Remove an Project from a Person
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /projects/{projectId} | Delete a project
[**GetOrganizationProjects**](ProjectApi.md#GetOrganizationProjects) | **Get** /organizations/{organizationId}/projects | Get a list of all Projects for an Organization
[**GetProject**](ProjectApi.md#GetProject) | **Get** /projects/{projectId} | Get details about a Project
[**GetProjects**](ProjectApi.md#GetProjects) | **Get** /projects | Get a list of all Projects in all Organizations
[**UpdatePersonProject**](ProjectApi.md#UpdatePersonProject) | **Put** /persons/{personId}/projects/{projectId} | Update a Project of a Person
[**UpdateProject**](ProjectApi.md#UpdateProject) | **Put** /projects/{projectId} | Update a Project



## AddPersonProject

> PersonDetails AddPersonProject(ctx, personId, projectId).SkillLevel(skillLevel).Execute()

Add Project to a Person

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    personId := TODO // string | 
    projectId := TODO // string | 
    skillLevel := []openapiclient.SkillLevel{*openapiclient.NewSkillLevel()} // []SkillLevel | A list of Skills each with a level

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.AddPersonProject(context.Background(), personId, projectId).SkillLevel(skillLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.AddPersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonProject`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.AddPersonProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**projectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skillLevel** | [**[]SkillLevel**](SkillLevel.md) | A list of Skills each with a level | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmSkill

> PersonDetails ConfirmSkill(ctx, personId, projectId, skillId, confirmingPersonId).Execute()

Confirm Skill

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    personId := TODO // string | 
    projectId := TODO // string | 
    skillId := TODO // string | 
    confirmingPersonId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.ConfirmSkill(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ConfirmSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ConfirmSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**projectId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 
**confirmingPersonId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> OrganizationDetails CreateProject(ctx, organizationId).Project(project).Execute()

Create a Project in an Organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := TODO // string | 
    project := *openapiclient.NewProject("Id_example", "This is the name") // Project | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.CreateProject(context.Background(), organizationId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**Project**](Project.md) |  | 

### Return type

[**OrganizationDetails**](OrganizationDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfirmation

> Status DeleteConfirmation(ctx, personId, projectId, skillId, confirmingPersonId).Execute()

Remove a confirmation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    personId := TODO // string | 
    projectId := TODO // string | 
    skillId := TODO // string | 
    confirmingPersonId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.DeleteConfirmation(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeleteConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfirmation`: Status
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.DeleteConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**projectId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 
**confirmingPersonId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Status**](Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePersonProject

> Status DeletePersonProject(ctx, personId, projectId).Execute()

Remove an Project from a Person

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    personId := TODO // string | 
    projectId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.DeletePersonProject(context.Background(), personId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeletePersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonProject`: Status
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.DeletePersonProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**projectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Status**](Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> Status DeleteProject(ctx, projectId).Execute()

Delete a project

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.DeleteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProject`: Status
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.DeleteProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationProjects

> PagedProjects GetOrganizationProjects(ctx, organizationId).Execute()

Get a list of all Projects for an Organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.GetOrganizationProjects(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetOrganizationProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationProjects`: PagedProjects
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetOrganizationProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PagedProjects**](PagedProjects.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> ProjectDetails GetProject(ctx, projectId).Execute()

Get details about a Project

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.GetProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectDetails**](ProjectDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> PagedProjects GetProjects(ctx).Skip(skip).Limit(limit).Execute()

Get a list of all Projects in all Organizations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    skip := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.GetProjects(context.Background()).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjects`: PagedProjects
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedProjects**](PagedProjects.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePersonProject

> PersonDetails UpdatePersonProject(ctx, personId, projectId).SkillLevel(skillLevel).Execute()

Update a Project of a Person

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    personId := TODO // string | 
    projectId := TODO // string | 
    skillLevel := []openapiclient.SkillLevel{*openapiclient.NewSkillLevel()} // []SkillLevel | A list of Skills each with a level

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.UpdatePersonProject(context.Background(), personId, projectId).SkillLevel(skillLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.UpdatePersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonProject`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.UpdatePersonProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**projectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skillLevel** | [**[]SkillLevel**](SkillLevel.md) | A list of Skills each with a level | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> ProjectDetails UpdateProject(ctx, projectId).Project(project).Execute()

Update a Project

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := TODO // string | 
    project := *openapiclient.NewProject("Id_example", "This is the name") // Project | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.UpdateProject(context.Background(), projectId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**Project**](Project.md) |  | 

### Return type

[**ProjectDetails**](ProjectDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
