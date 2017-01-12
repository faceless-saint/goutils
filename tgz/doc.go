/*
 * goutils/tgz
 *
 * Package tgz provides a convenient interface for working with gzipped
 * tar files. The Reader and Writer types have the exact same methods as
 * their conterparts in the archive/tar package, but internally they use
 * gzip for transparent compression and decompression of the data.
 *
 * This package allows a user working with compressed archives to access
 * the archive contents as a single step, rather than needing to manually
 * manage the individual artifacts from each package. Additionally, new
 * convenience methods are defined for common archive operations.
 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package tgz
