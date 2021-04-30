# passgen v1.0.0
Copyright 2021 arataca89@gmail.com<br>
Licensed under the MIT License<br>
20210429<br>
passgen é um simples gerador de senhas.
<pre>
USO:
passgen
        Exibe o help.

passgen char
        Exibe o conjunto de caractres ASCII utilizados para gerar a senha.

passgen c nr
        Cria uma senha de nr caracteres.
        Caso nr não seja definido cria uma senha de 16 caracteres.

A senha é criada e copiada para o clipboard.
Use CTRL + V para colar a senha onde você quiser.
</pre>
## TESTES
<pre>
C:\Users\nerd\golang\go_estudo\passgen>passgen c
V3#EY[2jI$E*RWgt

C:\Users\nerd\golang\go_estudo\passgen>passgen c 8
3UlZj5#N

C:\Users\nerd\golang\go_estudo\passgen>passgen c 10
.uE1;tu;>7

</pre>
