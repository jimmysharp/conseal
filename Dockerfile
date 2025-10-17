FROM scratch
COPY conseal /bin/conseal

EXPOSE 18212
ENTRYPOINT ["/bin/conseal"]