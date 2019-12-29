FROM scratch
COPY f5er /f5er
ENTRYPOINT ["/f5er"]
CMD [ "--help" ]
