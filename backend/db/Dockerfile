FROM postgres:17

ENV POSTGRES_DB bd2
ENV POSTGRES_USER root
ENV POSTGRES_PASSWORD root

EXPOSE 5432

COPY ./src/* ./
COPY . .

RUN chmod +x ./init_script.sh

CMD "./init_script.sh"

