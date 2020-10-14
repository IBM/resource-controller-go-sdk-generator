# resource-controller-go-sdk-generator
Code Generation Project to Create the Golang SDK for the "IBM Cloud Resource Controller API"

## Usage
To download (from command-line):
```
go get github.com/IBM/resource-controller-go-sdk-generator v1.0.0
```

Then import the generated client:
```
eventstreams "github.com/IBM/resource-controller-go-sdk-generator/build/generated"
```
See the generated [documentation](build/generated/README.md) for details about the client.

## Code Generation
###### NOTE: This SDK Generation project/Guide uses the OpenAPI Generator Gradle Plugin (version "4.3.0"). This is used for Open API 3.x specification (swagger) documents. 

### How to generate the Cloud Resource Controller Golang Client
- Validate the Swagger definition: `./gradlew openApiValidate`
- Generate the Go code via gradle task: `./gradlew openApiGenerate`

### How to update the Swagger definition
- Download the newest Swagger spec from https://cloud.ibm.com/apidocs/resource-controller and save it as `documentation/swagger.json`.
  - NOTE: see this Post  for how to download the "OpenAPI definition JSON document": https://stackoverflow.com/questions/57154077/where-can-i-find-the-swagger-documentation-json-doc-for-the-ibm-cloud-iam-iden
, generic Swagger Codegen system is unable to handle these tags.  
- In the file build.gradle, change the value of `majorMinorVersion` to reflect whatever changes were made in the Swagger definition since the last client generation (using semantic versioning).
- When ready to "publish" the new version, simply push the newly generated code changes. Then once it's in master, tag the master branch with "v\<majorMinorVersion\>", which will allow the new version to be accessible by other go code. For example, to publish the 1.0.2 version, you'd run the following git commands:
  - git tag v1.0.2
  - git push origin v1.0.2

## Contributions
Follow the instructions above to make changes to the API specification and regenerate the code. Make the changes in a fork and then submit a pull request.
