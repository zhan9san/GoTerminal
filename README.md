# GoTerminal

Test GoTerminal

1. Get the test code

    ```bash
    git clone https://github.com/zhan9san/GoTerminal.git
    ```

2. Build a binary, IsTerminal

    ```bash
    cd GoTerminal/IsTerminal
    go build
    ```

3. Verify the difference

    1. Run in terminal

        ```bash
        $ ./IsTerminal
        IsTerminal:  true
        ```

    2. Run IsTerminal thru Golang os/exec

        ```bash
        $ cd ../ExecIsTerminal
        $ go run main.go
        Starting command
        out: IsTerminal:  false

        err:
        ```
