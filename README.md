# dockergetnetworkip
Gets Docker container's IP based on a network name.

#Usage
```
./dockergetnetworkip {DOCKERNETWORKNAME}
````

You should include the Docker Socket on the Container
```
docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock ubuntu ./dockergetnetworkip bridge
```
