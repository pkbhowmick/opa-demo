package example

default hello = false
default code = false
default no_public_network = false
default no_slow_query = false

hello {
    m := input.message
    m == "hello"
}

code {
    m := input.statusCode
    m == 200
}

no_public_network {
 not any_public_network
}

any_public_network {
    input.networks[_].public == true
}

no_slow_query {
 not any_slow_query
}

any_slow_query {
    input.dbQuery[_].execTime >= 10.0
}