### Forge

### Commands

- ```forge init```
- ```forge add```
- ```forge search```
- ```forge build```

### Other benefits

- Uses toml to build a pom.xml
  - do we want to make it reverse comatable so toml -> pom AND pom -> toml

### Outstanding questions

- Lock file?
  - Do we want to use the pom.xml at all or have it generate on the fly and not be commited???
- Other Maven features
  - profiles
  - plugins
  - repos
  - multi module projects

### TODOS

- Search Functionality [More Info](Search.md)
  - Proper searching of maven central (will be a nightmare)

- Add Functionality [More Info](Add.md)
  - building the forge.toml

- Build Functionality [More Info](Build.md)
  - building the forge.lock + pom.xml
