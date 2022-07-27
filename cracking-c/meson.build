project('leetcode_c', 
  'cpp',
  default_options : ['warning_level=3', 'cpp_std=c++20']
)

gtest_dep = dependency('gtest', fallback : ['gtest', 'gtest_dep'])

exe2_1    = executable('2.1.remove_dups', 'src/2.1.remove_dups.cc', cpp_args : '-DWITHOUT_BUFFER')

exe2_2    = executable('2.2.k_th_to_last',
 'src/2.2.k_th_to_last.cc',
  cpp_args : '-DLEETCODE_LIST',
  dependencies: [gtest_dep],
  install: true
)

test('basic', exe2_2)
