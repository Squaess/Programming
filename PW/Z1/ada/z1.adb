with Ada.Text_IO; use Ada.Text_IO;
with Ada.Strings.Unbounded;


procedure z1 is 
    File : File_Type;
    Value : Ada.Strings.Unbounded.Unbounded_String;

begin
    Open(File => File,
         Mode => In_File,
         Name => "../data.txt");
    loop
        exit when End_Of_File(File);
        Value := Ada.Strings.Unbounded.To_Unbounded_String(Get_Line(File));
        Put(Ada.Strings.Unbounded.To_String(Value));
        New_Line;
    end loop;

    Close(File);
end z1;
