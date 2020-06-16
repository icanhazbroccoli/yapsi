program procedurecall;

var
    a : integer;
    b : integer;

procedure square(n : integer);
begin
    n := n * n
end

procedure pow(n, p : integer);
var
    i : integer;
    n0 : integer;
begin
    writeln(n, p);
    i := 0;
    n0 := n;
    while i < p - 1 do
    begin
        writeln(n);
        n := n * n0;
        i := i + 1
    end
end

begin
    a := 123;
    square(a);
    writeln(a);

    b := 2;
    pow(b, 10);
    writeln(b)
end.
