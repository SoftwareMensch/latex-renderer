# LaTeX Renderer

Microservice which is responsible for LaTeX based PDF rendering.

# Start docker environment

```
sudo docker-compose up
```

# Render your file
```
curl -X POST \
     -o document.pdf \
     -H "Content-Type: multipart/form-data" \
     -F "document=@document.tex" \
     http://172.99.1.1:80/api/v1/document
```

