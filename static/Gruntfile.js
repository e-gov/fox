// Generated on 2014-05-09 using generator-angular 0.8.0
'use strict';

// # Globbing
// for performance reasons we're only matching one level down:
// 'test/spec/{,*/}*.js'
// use this if you want to recursively match all subfolders:
// 'test/spec/**/*.js'

module.exports = function (grunt) {

    require('load-grunt-tasks')(grunt);


    var proxySnippet = require('grunt-connect-proxy/lib/utils').proxyRequest;

    var props = grunt.file.readJSON('properties.json');

    grunt.initConfig({
        yeoman: {
            app: 'app',
            dist: 'dist'
        },
        // Watches files for changes and runs tasks based on the changed files
        watch: {
            js: {
                files: ['<%= yeoman.app %>/**/{,*/}*.js'],
                tasks: ['newer:jshint:all'],
                options: {
                    livereload: '<%= connect.options.livereload %>'
                }
            },
            sass: {
                files: ['<%= yeoman.app %>/styles/{,*/}*.{scss,sass}'],
                tasks: ['sass:server', 'autoprefixer']
            },
            gruntfile: {
                files: ['Gruntfile.js']
            },
            livereload: {
                options: {
                    livereload: '<%= connect.options.livereload %>'
                },
                files: [
                    '<%= yeoman.app %>/{,*/}*.html',
                    '<%= yeoman.app %>/{,**/}*.html',
                    '.tmp/css/{,*/}*.css',
                    '<%= yeoman.app %>/{,*/}*.{png,jpg,jpeg,gif,webp,svg}'
                ]
            }
        },
        connect: {
            //server: {
            options: {
                port: 9000,
                hostname: '0.0.0.0',
                keepalive: true,
                open: true,
                livereload: 35729,
                base: ['<%= yeoman.app %>'],
                middleware: function (connect, options, middlewares) {
                    middlewares.push(proxySnippet);
                    return middlewares;
                }
            },
            proxies: props.proxies,
            livereload: {
                options: {
                    open: false,
                    base: [
                        '.tmp',
                        '<%= yeoman.app %>'
                    ]
                }
            },
            dist: {
                options: {
                    base: '<%= yeoman.dist %>'
                }
            }
            //}
        },
        clean: {
            dist: {
                files: [
                    {
                        dot: true,
                        src: [
                            '.tmp',
                            '<%= yeoman.dist %>/*',
                            '!<%= yeoman.dist %>/.git*'
                        ]
                    }
                ]
            },
            distModules: ['<%= yeoman.dist %>/modules'],
            server: '.tmp'
        },
        sass: {
            options: {
                includePaths: [
                'bower_components'
                ]
            },
            dist: {
                files: [{
                    expand: true,
                    cwd: '<%= yeoman.app %>/styles',
                    src: ['*.scss'],
                    dest: '.tmp/styles',
                    ext: '.css'
                }]
            },
            server: {
                files: [{
                    expand: true,
                    cwd: '<%= yeoman.app %>/styles',
                    src: ['*.scss'],
                    dest: '.tmp/styles',
                    ext: '.css'
                }]
            }
        },
        autoprefixer: {
            options: {
                browsers: ['last 1 version']
            },
            dist: {
                files: [
                    {
                        expand: true,
                        cwd: '.tmp/css/',
                        src: '{,*/}*.css',
                        dest: '.tmp/css/'
                    }
                ]
            }
        },
        copy: {
            dist: {
                files: [
                    {
                        expand: true,
                        dot: true,
                        cwd: '<%= yeoman.app %>',
                        dest: '<%= yeoman.dist %>',
                        src: [
                            '*.{ico,png,txt}',
                            '.htaccess',
                            '*.html',
                            'views/{,*/}*.html',
                            'images/{,*/}*.{webp}',
                            'fonts/*'
                        ]
                    },
                    {
                        expand: true,
                        cwd: '.tmp/images',
                        dest: '<%= yeoman.dist %>/images',
                        src: ['generated/*']
                    },
                    {
                        expand: true,
                        cwd: '<%= yeoman.app %>',
                        dest: '<%= yeoman.dist %>',
                        src: ['bower_components/{,*/}*/fonts/*', 'bower_components/angular-chosen-localytics/spinner.gif', 'bower_components/chosen/chosen-sprite.png']
                    }
                ]
            },
            styles: {
                expand: true,
                cwd: '<%= yeoman.app %>/css',
                dest: '.tmp/css/',
                src: '{,*/}*.css'
            }

        },
        concurrent: {
            server: [
                'sass:server',
                'copy:styles'
            ]
        }
    });

    grunt.loadNpmTasks('grunt-connect-proxy');
    grunt.loadNpmTasks('grunt-contrib-connect');

    //grunt.registerTask('serve', [
    //    'configureProxies:server',
    //    'connect:server'
    //]);


    grunt.registerTask('serve', [
        'clean:server',
        'concurrent:server',
        'autoprefixer',
        'configureProxies:server',
        'connect:livereload',
        'watch'
    ]);

};
