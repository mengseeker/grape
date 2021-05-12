FROM envoybase

ENV ETCD_VER=v3.4.16
ENV GOOGLE_URL=https://storage.googleapis.com/etcd
ENV GITHUB_URL=https://github.com/etcd-io/etcd/releases/download

RUN DOWNLOAD_URL=${GOOGLE_URL} &&\
  rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz &&\
  rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test &&\
  curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz &&\
  tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1 &&\
  cp /tmp/etcd-download-test/etcd /usr/local/bin/ &&\
  cp /tmp/etcd-download-test/etcdctl /usr/local/bin/ &&\
  rm -rf /tmp/etcd-download-test  &&\
  etcd --version &&\
  etcdctl version

CMD ["etcd"]