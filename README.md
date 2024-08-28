# falco-rules-gen


Falco-rules-gen is an assistant for generating custom Falco rules, macros lists, among other things.


Development enviroment execution

    go run .


Binary installation

    go install ./...  (Locally)

    go install github.com/diego-all/falco-rules-gen@latest


Uninstall the binary

    rm $(go env GOPATH)/bin/falco-rules-gen

    rm $(go env GOBIN)/falco-rules-gen


Execute the application

    falco-rules-gen generate my-custom-rule -o ./my-rule.yaml
