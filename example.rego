package example

default hello = false
default code = false

hello {
    m := input.message
    m == "hello"
}

code {
    m := input.statusCode
    m == 200
}
