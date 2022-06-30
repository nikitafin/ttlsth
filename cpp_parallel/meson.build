project('parallel',
  'cpp',
  version: '0.1.0',
  license: 'BSD',
  default_options : ['warning_level=3', 'cpp_std=c++20']
)

compiler     = meson.get_compiler('cpp')
compiler_id = compiler.get_id()

# FIXME: cant link with libc++
assert(compiler_id == 'clang')
compiler.find_library('libc++', required: true)
add_project_arguments (
  [
    '-Wmissing-declarations',
    '-Wno-unused',
    '-stdlib=libc++'
  ],

  language: 'cpp'
)
add_project_link_arguments (
  [
    '-lc++',
    '-lc++abi'
  ],

  language: 'cpp'
)


main = executable(
  'main',
  'src/main.cc',
)