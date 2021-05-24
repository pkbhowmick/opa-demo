package authz

test_allow_with_own_salary_request {
    allow with input as {"path":["salary", "alice"], "method":"GET", "user":"alice"}
}

test_deny_with_others_salary_request {
    not allow with input as {"user": "alice", "path": ["salary", "bob"], "method" : "GET"}
}

test_allow_manager_with_his_subordinates_salary_request {
    allow with input as {"user": "bob", "path": ["salary","alice"], "method": "GET"}
}

test_deny_manager_with_others_subordinates_salary_request {
    not allow with input as {"user" : "bob", "path": ["salary","charlie"], "method": "GET"}
}

test_allow_hr_manager_to_request_anyone_salary {
    allow with input as {"user" : "david", "path": ["salary", "bob"], "method": "GET"}
}
