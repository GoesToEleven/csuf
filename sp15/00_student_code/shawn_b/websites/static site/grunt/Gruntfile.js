module.exports = function (grunt) {
    grunt.initConfig({
        autoprefixer: {
            dist: {
                files: {
                    '../public/assets/build/main.css': '../public/assets/stylesheets/main.css'
                }
            }
        },
        watch: {
            styles: {
                files: ['main.css'],
                tasks: ['autoprefixer']
            }
        }
    });
    grunt.loadNpmTasks('grunt-autoprefixer');
    grunt.loadNpmTasks('grunt-contrib-watch');
    grunt.registerTask('default', ['autoprefixer']);
};