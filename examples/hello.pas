program HelloWorld;

function fib(n : integer) : integer;
begin
    if n <= 1
    then
        result := n
    else
        result := fib(n-1) + fib(n-2)
end;

begin
    writeln(fib(10))
end.
