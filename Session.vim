let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/git/type_writer
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +1 ~/git/type_writer
badd +69 front_end/src/components/game/Area.vue
badd +0 term://~/git/type_writer//89328:/bin/bash
argglobal
%argdel
$argadd ~/git/type_writer
edit NetrwTreeListing
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
wincmd _ | wincmd |
split
1wincmd k
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 45 + 123) / 246)
exe '2resize ' . ((&lines * 41 + 30) / 60)
exe 'vert 2resize ' . ((&columns * 200 + 123) / 246)
exe '3resize ' . ((&lines * 16 + 30) / 60)
exe 'vert 3resize ' . ((&columns * 200 + 123) / 246)
argglobal
balt ~/git/type_writer
setlocal foldmethod=expr
setlocal foldexpr=nvim_treesitter#foldexpr()
setlocal foldmarker={{{,}}}
setlocal foldignore=#
setlocal foldlevel=0
setlocal foldminlines=1
setlocal foldnestmax=20
setlocal nofoldenable
let s:l = 35 - ((34 * winheight(0) + 29) / 58)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 35
normal! 0
lcd ~/git/type_writer
wincmd w
argglobal
if bufexists(fnamemodify("~/git/type_writer/front_end/src/components/game/Area.vue", ":p")) | buffer ~/git/type_writer/front_end/src/components/game/Area.vue | else | edit ~/git/type_writer/front_end/src/components/game/Area.vue | endif
if &buftype ==# 'terminal'
  silent file ~/git/type_writer/front_end/src/components/game/Area.vue
endif
setlocal foldmethod=expr
setlocal foldexpr=nvim_treesitter#foldexpr()
setlocal foldmarker={{{,}}}
setlocal foldignore=#
setlocal foldlevel=0
setlocal foldminlines=1
setlocal foldnestmax=20
setlocal nofoldenable
let s:l = 57 - ((26 * winheight(0) + 20) / 41)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 57
normal! 056|
lcd ~/git/type_writer
wincmd w
argglobal
if bufexists(fnamemodify("term://~/git/type_writer//89328:/bin/bash", ":p")) | buffer term://~/git/type_writer//89328:/bin/bash | else | edit term://~/git/type_writer//89328:/bin/bash | endif
if &buftype ==# 'terminal'
  silent file term://~/git/type_writer//89328:/bin/bash
endif
balt ~/git/type_writer/front_end/src/components/game/Area.vue
setlocal foldmethod=expr
setlocal foldexpr=nvim_treesitter#foldexpr()
setlocal foldmarker={{{,}}}
setlocal foldignore=#
setlocal foldlevel=0
setlocal foldminlines=1
setlocal foldnestmax=20
setlocal nofoldenable
let s:l = 15 - ((6 * winheight(0) + 8) / 16)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 15
normal! 0
lcd ~/git/type_writer
wincmd w
2wincmd w
exe 'vert 1resize ' . ((&columns * 45 + 123) / 246)
exe '2resize ' . ((&lines * 41 + 30) / 60)
exe 'vert 2resize ' . ((&columns * 200 + 123) / 246)
exe '3resize ' . ((&lines * 16 + 30) / 60)
exe 'vert 3resize ' . ((&columns * 200 + 123) / 246)
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
nohlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
