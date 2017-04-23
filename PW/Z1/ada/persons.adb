with dates; use dates; -- definition of type date

                                            -- Chapter 22 - Program 2
package Persons is

   -- This is a very simple transportation type.
   type person is
      record
         Name : string(1..20);
         Surname : string(1..20);
      end record;

   procedure Set_Values(Person_In : in out person;
                        Name  : Integer;
                        Surname  : string(1..20));
   function Get_Name(Person_In : person) return string(1..20);
   function Get_Surname(Person_In : person) return string(1..20);

   -- This CAR type extends the functionality of the TRANSPORT type.

end Persons;




package body Persons is

-- Subprograms for the TRANSPORT record type.
procedure Set_Values(Person_In : in out person;
                     Name  : string(1..20);
                     Surname  : string(1..20)) is
begin
   Person_In.Name := Name;
   Person_In.Surname := Surname;
end Set_Values;

function Get_Name(Person_In : person) return string(1..20) is
begin
   return Person_In.Name;
end Get_Name;

function Get_Surname(Person_In : person) return string(1..20) is
begin
   return Person_In.Surname;
end Get_Surname;

end Persins;
