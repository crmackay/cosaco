# how we will deploy...

1. verify that travisCI passes

1. send a githook to digital ocean web server

1. webserver handles the githook and conducts a git pull

1. webserver builds the new program
    - if build errors or anything log it and stop (and send me an email?)

1. webserver gracefully transfers requests to the new server...
    - how to do this?
        - use the endless package?
        - roll own golang solution?
        - use nginx load balancing to switch from one port to another
            - then close and stop old server
