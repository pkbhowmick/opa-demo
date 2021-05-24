package authz

managers = {"alice": [], "charlie": [], "bob":["alice"], "betty": ["charlie"]}
hr = ["david"]

default allow = false

# Allow users to get their own salary
allow {
    some username
    input.path = ["salary", username]
    input.method == "GET"
    input.user = username
}

# Allow managers to get their employees' salary
allow {
    some username
    input.method == "GET"
    input.path = ["salary", username]
    managers[input.user][_] == username
}

# Allow HR manager to get anyone's salary
allow {
    input.method == "GET"
    input.path = ["salary", _]
    input.user == hr[_]
}
