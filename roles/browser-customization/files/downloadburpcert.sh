#!/bin/bash
bash -c "timeout 45 /usr/lib/jvm/jdk-21*/bin/java -Djava.awt.headless=true -jar /opt/burpsuite.jar < <(echo y) &"
sleep 30
curl localhost:8080/cert -o /tmp/cacert.der
exit
