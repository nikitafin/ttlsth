project('in_discovering_modem_cpp',
  'cpp',
  version: '0.1.0',
  license: 'BSD',
  default_options : ['warning_level=3', 'cpp_std=c++20']
)

compiler    = meson.get_compiler('cpp')
compiler_id = compiler.get_id()

compiler_flags = []
linker_flags   = []

if compiler_id == 'clang'
  compiler_flags += [
    '-Wmissing-declarations',
    '-Wno-unused'
  ]

  if compiler.find_library('libc++', required: false).found()
    compiler_flags += [
      '-stdlib=libc++'
    ]
    linker_flags += [
      '-lc++',
      '-lc++abi'
    ]
  endif

elif compiler_id ==  'msvc'
    compiler_flags += [
      '/Od',
      '/EHc',
      # '/Fa',
      '/Zi',
      '/MP',
      '/sdl'
    ]
else
  assert(false)
endif


add_project_arguments (
  compiler_flags,

  language: 'cpp'
)

add_project_link_arguments (
  linker_flags, 

  language: 'cpp'
)

age = executable(
  'age',
  'src/1.age.cc'
)