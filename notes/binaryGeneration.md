# Binary Generation



    go install ./...
    which falco-rules-gen


## Go BIN and Go PATH

go install no actualiza automáticamente el PATH del sistema. Lo que hace go install es compilar el código y colocar el binario resultante en el directorio especificado en la variable GOBIN. Si GOBIN no está configurado, el binario se coloca en el directorio GOPATH/bin.

Para asegurarte de que el binario instalado esté disponible en tu PATH, debes:
Verificar el directorio GOPATH/bin:

Por defecto, GOPATH se establece en $HOME/go. Entonces, los binarios se colocan en $HOME/go/bin.
Puedes comprobar tu GOPATH ejecutando:
bash

    go env GOPATH

El binario se colocará en $GOPATH/bin.

Añadir GOPATH/bin al PATH:

Si GOPATH/bin no está en tu PATH, puedes añadirlo temporalmente:


    export PATH=$PATH:$(go env GOPATH)/bin

Para hacerlo permanente, añade la línea anterior a tu archivo de configuración de shell, como .bashrc, .zshrc, o .bash_profile dependiendo del shell que uses:

    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
    source ~/.bashrc

Verificar si GOBIN está configurado:

Si has configurado la variable GOBIN, los binarios se instalarán en esa ubicación en lugar de GOPATH/bin.
Puedes verificar si GOBIN está configurado y dónde apunta ejecutando:

    go env GOBIN

Si GOBIN está configurado, asegúrate de que su valor esté en tu PATH.

Resumen:

go install instala el binario en $GOBIN o, si no está configurado, en $GOPATH/bin.
No actualiza automáticamente tu PATH. Debes asegurarte de que $GOPATH/bin o $GOBIN estén en tu PATH para que puedas ejecutar el binario desde cualquier lugar en la terminal.



