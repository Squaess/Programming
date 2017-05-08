package body Conveyance3 is

-- Subprograms for the TRANSPORT3 record
procedure Set_Values(Vehicle_In : in out TRANSPORT; 
                     Wheels_In : INTEGER; 
                     Weight_In : FLOAT) is
begin
   Vehicle_In.Wheels := Wheels_In;
   Vehicle_In.Weight := Weight_In;
end Set_Values;


function Get_Wheels(Vehicle_In : TRANSPORT) return INTEGER is
begin
   return Vehicle_In.Wheels;
end Get_Wheels;


function Get_Weight(Vehicle_In : TRANSPORT) return FLOAT is
begin
   return Vehicle_In.Weight;
end Get_Weight;


-- Subprograms for the CAR record
procedure Set_Values(Vehicle_In : in out CAR; 
                     Passenger_Count_In : INTEGER) is
begin
   Vehicle_In.Passenger_Count := Passenger_Count_In;
end Set_Values;


function Get_Passenger_Count(Vehicle_In : CAR) return INTEGER is
begin
   return Vehicle_In.Passenger_Count;
end Get_Passenger_Count;


-- Subprograms for the TRUCK record
procedure Set_Values(Vehicle_In : in out TRUCK; 
                     Wheels_In : INTEGER; 
                     Weight_In : FLOAT; 
                     Passenger_Count_In : INTEGER; 
                     Payload_In : FLOAT) is
begin
      -- This is one way to set the values in the base class
   Vehicle_In.Wheels := Wheels_In;
   Vehicle_In.Weight := Weight_In;

      -- This is another way to set the values in the base class
   Set_Values(TRANSPORT(Vehicle_In), Wheels_In, Weight_In);

      -- This sets the values in this class
   Vehicle_In.Passenger_Count := Passenger_Count_In;
   Vehicle_In.Payload := Payload_In;
end Set_Values;


function Get_Passenger_Count(Vehicle_In : TRUCK) return INTEGER is
begin
   return Vehicle_In.Passenger_Count;
end Get_Passenger_Count;

end Conveyance3;



-- Result of execution
--
-- (This program cannot be executed alone.)

