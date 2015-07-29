# cosaco

![travisCI](https://travis-ci.org/crmackay/cosaco.svg?branch=master) ![travisCI](https://travis-ci.org/crmackay/cosaco.svg?branch=master) [![Coverage Status](https://coveralls.io/repos/crmackay/cosaco/badge.svg?branch=master&service=github)](https://coveralls.io/github/crmackay/cosaco?branch=master)

webapp for the COSACO initiative of La Romana, Dominican Republic (el sitio del la colaboración de COSACO de La Romana, Republica Dominicana)


## main goals and technolgies:

- to start this will be a [`golang`](http://golang.org) driven web app, which will allow future expansion into an API backend which can serve several frontends (mobile apps, web apps, etc.) for data-sharing and collaboration functionalities later on
- this whole thing will be rolled out in a interative way, so from step 1 we will have zero downtime upgrades built in
- as more features get added the binary will be seamlessly updated

- see ideas [here](https://docs.google.com/document/d/12b7MaSla2uKhMK9sewEevJHai1DO84s9vq3e-SUWU1c/edit?usp=sharing)

### features:

- version 0.1:
  - graceful restart (when new binaries are loaded to the server)
    - set up githook that pulls, builds, find the PID for the current server and kills it
    - use [endless](https://github.com/fvbock/endless) to gracefully reload
  - homepage
  	- pages are in a buffer
  - data stored in a json files
  	- page buffers update when json files update
  	- [use fsnotify](https://godoc.org/gopkg.in/fsnotify.v1)
  - spanish and English toggle
  - i18n (Spanish and English, with translator-editable files)
  - members page
    - a page for each member, with details, contact info etc.

- version 0.2
	- database instead of data instead of json
	- notifyer...on update push to pages
   
