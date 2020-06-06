program HelloWorld;

var
    i : integer;
    foo : real

begin

    i := 0;
    foo := -12.345;

    while i < 10 do
    begin
        writeln('Hello World!');
        i := i + 1
    end

end.
