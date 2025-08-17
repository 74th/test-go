ls
cd /root
cat <<EOF > main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
EOF

ls
