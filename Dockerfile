FROM heroku/heroku:16-build as build

COPY . /app
WORKDIR /app

# Setup buildpack
RUN make
RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
RUN curl https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go

#Execute Buildpack
RUN STACK=heroku-16 /tmp/buildpack/heroku/go/bin/compile /bin /tmp/build_cache /tmp/env

# Prepare final, minimal image
FROM heroku/heroku:16

COPY --from=build /app /app
ENV HOME /app
WORKDIR /app
RUN useradd -m heroku
USER heroku
CMD /app/opt