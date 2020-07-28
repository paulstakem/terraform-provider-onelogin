resource onelogin_saml_apps saml{
  connector_id = 50534
  name =  "Updated SAML App"
  description = "Updated SAML"

  configuration = {
    signature_algorithm = "SHA-256"
  }
}

resource onelogin_app_rules test_a{
  enabled = true
  match = "all"
  name = "updated first rule"
  app_id = onelogin_saml_apps.saml.id
  conditions {
    operator = ">"
    source = "last_login"
    value = "90"
  }
  actions {
    action = "set_amazonusername"
    expression = ".*"
    value = ["member_of"]
  }
}

resource onelogin_app_rules test_b{
  enabled = true
  match = "all"
  name = "updated second rule"
  app_id = onelogin_saml_apps.saml.id
  conditions {
    operator = "ri"
    source = "has_role"
    value = "340475"
  }
  actions {
    action = "set_amazonusername"
    expression = ".*"
    value = ["member_of"]
  }
  depends_on = [onelogin_app_rules.test_a]
}
