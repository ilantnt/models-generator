# Models-Generator
The primary goal of this project is to automate the creation of Go structs based on YAML models. By expressing data structures in a declarative YAML format, developers can harness the simplicity of the Go language without engaging in the manual, error-prone process of writing repetitive code. A key feature is the introduction of a make generate command, which transforms YAML models into Go structs, effectively streamlining the development process.



## Features
* Automated Struct Generation: Express data models in YAML, and utilize the make generate command to automatically generate corresponding Go structs.

* Synchronization with TypeScript Definitions: Recognizing the significance of collaboration between frontend and backend teams, the project introduces a make sync command. This command goes beyond code generation, synchronizing the generated Go structs with TypeScript definitions.

* Bidirectional Synchronization: Ensure changes to data models are reflected in both backend Go code and frontend TypeScript code without manual intervention, fostering a seamless development environment.

## Usage

Define YAML Models
```yaml
models:
  Student:
    Classes: 
        type: "list"
        itemType: "string"
    Grade:
        type: "number"
  Person:
        Name:
            type: "string"
        Age:
            type: "number"
        Addresses:
            type: "list"
            itemType: "Address"
  Address:
        City: 
            type: "string"
        Country:
            type: "string"
```
Then Generate Go Structs:

```makefile
make generate
```
To synchronize the generated Go structs with TypeScript definitions, use the following command:

```makefile
make sync
```
## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.
