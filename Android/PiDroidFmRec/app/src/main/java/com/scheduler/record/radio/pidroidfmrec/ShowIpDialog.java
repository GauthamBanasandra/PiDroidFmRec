package com.scheduler.record.radio.pidroidfmrec;

import android.app.AlertDialog;
import android.app.Dialog;
import android.app.DialogFragment;
import android.content.Context;
import android.content.DialogInterface;
import android.net.wifi.WifiManager;
import android.os.Bundle;
import android.text.format.Formatter;

/**
 * Created by gauth_000 on 04-Dec-16.
 */
public class ShowIpDialog extends DialogFragment
{
    @Override
    public Dialog onCreateDialog(Bundle savedInstanceState)
    {

        AlertDialog.Builder builder = new AlertDialog.Builder(getActivity());
        builder.setMessage(getIp())
                .setTitle("Your device IP")
                .setPositiveButton("Ok", new DialogInterface.OnClickListener()
                {
                    @Override
                    public void onClick(DialogInterface dialog, int which)
                    {

                    }
                });

        return builder.create();
    }

    private String getIp()
    {
        WifiManager wm = (WifiManager) getContext().getSystemService(Context.WIFI_SERVICE);
        return Formatter.formatIpAddress(wm.getConnectionInfo().getIpAddress());
    }
}
