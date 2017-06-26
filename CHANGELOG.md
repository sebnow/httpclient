# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2017-06-26
### Added
- The `Client` interface adhering to the `http.Client` concrete
  implementation.
- The `ClientContext` interface and wrapper implementation, which adds
  `context.Context` support to the Get/Post/Head/PostForm convenience
  methods.
