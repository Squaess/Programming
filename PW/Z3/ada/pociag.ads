--limited with model;
with Ada.Containers.Vector;
with Ada.Strings.Unbounded;
with Ada.Real_Time;
--with pracownik; use pracownik;

package pociag is 
    type Pociag_Typ is (Pociag_Typ_N, Pociag_Typ_S);
    type POCIAG (p_typ:Pociag_Typ);

    task type PociagTask ( pociag_ptr : access POCIAG;
                           model_ptr : access model.Symulacja_Model) is
        entry koniecZwrotnica(zwrotnica_id : in Positive);
        entry koniecPeron(peron_id : in Positive);
        entry koniecTrasy(
