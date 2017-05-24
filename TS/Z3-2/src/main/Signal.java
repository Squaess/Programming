package main;

/**
 * Created by aedd on 5/10/17
 */

class Signal {

    private Host master;
    private int position;
    private boolean direction;
    private int signalDuration;
    private boolean inSource = true;

    Signal(Host master, int position, boolean direction, int signalDuration)
    {
        this.master = master;
        this.position = position;
        this.direction = direction;
        this.signalDuration = signalDuration;
    }

    void move()
    {
        signalDuration--;

        if(inSource)
        {
            inSource = false;
            return;
        }


        //Jeżeli sygnał "porusza się" w prawo
        if(direction)
        {
            position++;

            if(position >= Network.getLength())
            {
                signalDuration = 0;
                //position = Network.getLength()-1;
                //direction = false;
                position--;
            }
        }
        else
        {
            position--;
            if(position < 0)
            {
                signalDuration = 0;
                //position = 0;
                //direction = true;
                position++;
            }
        }
    }

    Host getMaster()
    {
        return master;
    }

    int getPosition()
    {
        return position;
    }

    int getSignalDuration()
    {
        return signalDuration;
    }

    boolean getDirection()
    {
        return direction;
    }

    void lose()
    {
        master = null;
    }
}
