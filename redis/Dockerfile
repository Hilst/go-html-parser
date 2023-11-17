FROM redis
COPY . /tmp
RUN chmod -R 777 /tmp
ENTRYPOINT ["/tmp/entrypoint.sh"]
CMD ["redis-server"] 