# goci

goci is a CLI used to update [openware/versions](github.com/openware/versions) in every component's CI pipeline

## Usage

To update a given component's version in `openware/versions`, run:
```sh
goci versions -path opendax/*branch* -component *component_name* -tag *tag*
```

This will do the following:
1. Clone `openware/versions`
2. Update `opendax/2-6/versions.yaml` component tag
3. Commit and push the updates

