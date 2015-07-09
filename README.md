# cosaco
webapp for the COSACO initiative of La Romana, Dominican Republic (el sitio del la colaboración de COSACO de La Romana, Republica Dominicana)


## main goals and technolgies:

- to start this will be a `go` driven web app, which will allow future expansion into an API backend which can serve several frontends (mobile apps, web apps, etc.)
- this whole thing will be rolled out in a interative way, so from step 1 we will have zero downtime upgrades built in
- as more features get added the binary will be seamlessly updated

### features:

- graceful restart (when new binaries are loaded to the server)
  - set up githook that pulls, builds, and reloads the binary
- i18n (Spanish and English, with translator-editable files)
- database of team info
- database of users and roles with access rules
