package com.scheduler.record.radio.pidroidfmrec;

import android.content.Context;
import android.content.Intent;
import android.support.v4.content.WakefulBroadcastReceiver;
import android.util.Log;

/**
 * Created by gauth_000 on 25-Nov-16.
 */
public class AlarmReceiver extends WakefulBroadcastReceiver
{
    private static final String TAG = "AlarmReceiver";

    @Override
    public void onReceive(Context context, Intent intent)
    {
        // Debug.
        Log.d(TAG, "Alarm received");

        startWakefulService(context, new Intent(context, WakeUpService.class));
    }
}
