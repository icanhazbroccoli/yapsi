program HelloWorld;

var
    i : integer;

function fib(n : integer) : integer;
begin
    if n <= 1
    then
        result := n
    else
        result := fib(n-1) + fib(n-2)
end;

begin

    i := 0;

    while i < 10 do
    begin
        writeln('iteration', i, ':', 'Hello World!');
        i := i + 1
    end;

    writeln('done')

end.
