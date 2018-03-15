FROM ruby:2.3

RUN mkdir /vehicles-ms
WORKDIR /vehicles-ms

ADD Gemfile /vehicles-ms/Gemfile
ADD Gemfile.lock /vehicles-ms/Gemfile.lock

RUN bundle install
ADD . /vehicles-ms