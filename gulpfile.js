var $ = require('gulp-load-plugins')();
var _ = require('lodash');
var browserify = require('browserify');
var buffer = require('vinyl-buffer');
var es = require('event-stream');
var exec = require('child_process').exec;
var git = require('git-rev');
var glob = require('glob');
var gulp = require('gulp');
var ngAnnotate = require('browserify-ngannotate');
var source = require('vinyl-source-stream');
var copiedFiles = require('./config/copiedFiles.json');
var packageJson = require('./package.json');
var cachebust = new $.cachebust();

var config = {
  outputName: 'app',
  templateCache: {
    file: "templates.js",
    options: {
      module: 'app',
      transformUrl: function(url) {
        return url.replace(/\.nghtml$/, ".html");
      }
    }
  },
  packedAssetName: 'assets.go',
  serveLiveAssets: !$.util.env.production,
  appName: packageJson.name,
  version: packageJson.version
};
var paths = {
  base: "assets",
  js: "assets/js/**.js",
  angularIncludes: 'assets/js/*/',
  scss: ["assets/css/**/*.scss", "assets/css/**/*.css", "!assets/css/colors.scss"],
  images: "assets/images/**/*",
  html: "assets/*.html",
  templates: "assets/**/*.nghtml",
  goTemplates: "templates/**.*",
  outputBase: "build",
  goOutput: "build/templates",
  assetOutput: "build/public",
  distOutput: "dist"
};

gulp.task('clean', function() {
  return gulp.src([
      paths.goOutput,
      paths.assetOutput,
      paths.distOutput,
      paths.outputBase + "/" + config.packedAssetName,
      paths.outputBase + "/" + config.outputName
    ]).
    pipe($.clean());
});

gulp.task("copy-files", function(cb) {
  var files = _.map(copiedFiles, function(fileGlobs, dest) {
    return gulp.src(fileGlobs).
      pipe(gulp.dest(dest)).
      pipe($.size({ title: "copied files" }));
  });
  es.merge(files).on('end', cb);
});

gulp.task("image-assets", function() {
  return gulp.src(paths.images).
    pipe($.imagemin({
      progressive: true,
      interlaced: true
    })).
    pipe(cachebust.resources()).
    pipe(gulp.dest(paths.assetOutput + "/images")).
    pipe($.size({ title: "images" }));
});

gulp.task("js-assets", function(cb) {
  glob(paths.js, function(err, files) {
    var bundles = _.map(files, function(filename) {
      return browserify({
          entries: filename,
          debug: true,
          paths: paths.angularIncludes,
          external: ['angular', 'jquery'],
          transform: [ngAnnotate]
        }).bundle().
        pipe(source(filename.replace(paths.base + '/', ''))).
        pipe(buffer()).
        pipe($.sourcemaps.init({ loadMaps: true })).
        pipe($.uglify()).
        pipe(cachebust.resources()).
        pipe($.sourcemaps.write('./')).
        pipe(gulp.dest(paths.assetOutput)).
        pipe($.size({ title: "script bundles" }));
    });
    es.merge(bundles).on('end', cb);
  });
});

gulp.task("style-assets", function() {
  return gulp.src(paths.scss).
    pipe($.sass().on('error', $.sass.logError)).
    pipe(cachebust.resources()).
    pipe(gulp.dest(paths.assetOutput + "/css")).
    pipe($.size({ title: "stylesheets" }));
});

gulp.task('templatecache', ["compile-assets"], function() {
  return gulp.src(paths.templates).
    pipe($.angularTemplatecache(config.templateCache.file, config.templateCache.options)).
    pipe(cachebust.resources()).
    pipe(cachebust.references()).
    pipe(gulp.dest(paths.assetOutput + "/js")).
    pipe($.size({ title: "templates" }));
});

gulp.task("html-pages", ["compile-assets", "templatecache"], function() {
  return gulp.src(paths.html).
    pipe(cachebust.references()).
    pipe(gulp.dest(paths.assetOutput)).
    pipe($.size({ title: "html pages"}));
});

gulp.task("go-templates:html", ["compile-assets", "templatecache"], function() {
  return gulp.src(paths.goTemplates).
    pipe(cachebust.references()).
    pipe(gulp.dest(paths.goOutput)).
    pipe($.size({ title: "Go template html pages"}));
});

gulp.task('bindata', ["build-web", "go-templates:html"], function(cb) {
  var extraParams = config.serveLiveAssets ? "-debug" : null;
  var cmd = [
    'go-bindata',
    '-o ' + config.packedAssetName,
    extraParams,
    '-prefix ' + paths.outputBase,
    paths.goOutput + '/...',
    paths.assetOutput + '/...'
  ].join(" ");
	exec(cmd, cb);
});

gulp.task('build-app', ['bindata'], function(cb) {
  git.long(function(buildSha) {
    var ldFlags = [
      "-X main.AppName='" + config.appName + "'",
      "-X main.BuildSHA='" + buildSha + "'",
      "-X main.Version='" + config.version + "'"
    ].join(' ');
    var cmd = [
      'go build',
      '-ldflags "' + ldFlags + '"',
      '-v',
      '-o ' + paths.distOutput + '/' + config.outputName,
      '.'
    ].join(' ');
    exec(cmd, cb);
  });
});

gulp.task("compile-assets", ["copy-files", "style-assets", "image-assets", "js-assets"]);
gulp.task("build-web", ["compile-assets", "templatecache", "html-pages"]);
gulp.task("build", ["build-app"]);

gulp.task("default", ["build-app"]);
