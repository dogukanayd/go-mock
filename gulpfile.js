gulp.task('coveralls', ['test'], function() { gulp.src('coverage/lcov.info').pipe(coveralls())});