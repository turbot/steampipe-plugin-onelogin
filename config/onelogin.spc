connection "onelogin" {
  plugin = "onelogin"

  # OneLogin Client ID for requests. Required.
  # See instructions at https://developers.onelogin.com/api-docs/1/getting-started/working-with-api-credentials
  # This can also be set via the ONELOGIN_CLIENT_ID environment variable
  # client_id = "22e2860f324de80ebae6b1ff1e2d8ce25f4e3460ac571d561487"

  # OneLogin Client Secret for requests. Required.
  # This can also be set via the ONELOGIN_CLIENT_SECRET environment variable
  # client_secret = "4bc85aa1acf895e90b816e15f3f3ec806ae8ce4635d9"

  # The OneLogin API uses your subdomain as the API Domain. This is the same address you would use to login to OneLogin. Required.
  # This can also be set via the ONELOGIN_SUBDOMAIN environment variable
  # subdomain = "mycompany"
}
