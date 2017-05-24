package main;

import java.util.Random;

/**
 * Created by aedd on 5/10/17
 */
class Host {
    private char message;
    private int position;
    private int signalDuration;
    private int transmissionDuration;

    private boolean isSending;

    private int errors;
    private int counter;
    private boolean error;

    Host(int position, char message, int signalDuration)
    {
        this.position = position;
        this.message = message;
        this.signalDuration = signalDuration;

        transmissionDuration = 0;

        isSending = false;

        error = false;
        errors = 0;
        counter = 0;
    }

    int getPosition()
    {
        return position;
    }

    char getMessage()
    {
        return message;
    }

    void start()
    {
        //System.out.print("\t" + message + " rozpoczyna transmisję.");

        isSending = true;
        this.transmissionDuration = Network.getLength()*2; //Network.getLength()*2
    }

    void error()
    {
        transmissionDuration = Network.getLength()*2;
        error = true;

        if(errors == 0)
        {
            errors = 2;
            counter = 2;
        }
        else if(errors < 1024)
        {
            errors *= 2;
            counter = new Random().nextInt(errors);
        }

        System.out.print("\t" + message + " został unieruchomiony (" + counter + ").");
    }

    void waitForTransmission()
    {
        if(counter > 0)
            counter--;
    }

    boolean isWaiting()
    {
        if(counter > 0)
        {
            isSending = true;
            return true;
        }

        return false;
    }

    boolean wantToTransmit()
    {
        if(transmissionDuration == 0)
        {
            if(isSending)
            {
                System.out.print("\t" + message + " udało się wysłac transmisję bez problemów.");
            }

            error = false;
            isSending = false;
        }

        return isSending;
    }

    Signal createSignal(boolean direction)
    {
        if(transmissionDuration == Network.getLength()*2)
        {
            if(error)
                System.out.print("\t" + message + " wznowił transmisję.");
            else
                System.out.print("\t" + message + " rozpoczyna transmisję.");
        }

        transmissionDuration--;

        return new Signal(this, position, direction, signalDuration);
    }
}
