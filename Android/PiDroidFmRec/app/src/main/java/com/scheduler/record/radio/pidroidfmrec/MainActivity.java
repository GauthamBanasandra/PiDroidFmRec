package com.scheduler.record.radio.pidroidfmrec;

import android.app.ActivityManager;
import android.app.AlarmManager;
import android.app.PendingIntent;
import android.app.admin.DevicePolicyManager;
import android.content.ComponentName;
import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.view.Menu;
import android.view.MenuInflater;
import android.view.MenuItem;
import android.view.View;
import android.view.Window;
import android.view.WindowManager;
import android.widget.Button;
import android.widget.TimePicker;
import android.widget.Toast;

import java.util.Calendar;

public class MainActivity extends AppCompatActivity
{
    private static final String TAG = "MainActivity";
    private ComponentName compName;

    @Override
    protected void onCreate(Bundle savedInstanceState)
    {
        super.onCreate(savedInstanceState);
        // Remove title bar.
        this.requestWindowFeature(Window.FEATURE_NO_TITLE);
        // Remove notification bar.
        this.getWindow().setFlags(WindowManager.LayoutParams.FLAG_FULLSCREEN, WindowManager.LayoutParams.FLAG_FULLSCREEN);
        setContentView(R.layout.activity_main);

        final TimePicker timePickerWakeUp = (TimePicker) findViewById(R.id.timePicker_wake_up);
        Button buttonSet = (Button) findViewById(R.id.button_set);

        // Getting handlers for managing device's power.
        DevicePolicyManager deviceManger = (DevicePolicyManager) getSystemService(Context.DEVICE_POLICY_SERVICE);
        ActivityManager activityManager = (ActivityManager) getSystemService(Context.ACTIVITY_SERVICE);
        compName = new ComponentName(this, AdminReceiver.class);

        // Launch the activity to get admin permission.
        get_admin_permission();

        assert buttonSet != null;
        buttonSet.setOnClickListener(new View.OnClickListener()
        {
            @Override
            public void onClick(View v)
            {
                assert timePickerWakeUp != null;
                Toast.makeText(MainActivity.this, timePickerWakeUp.getHour() + "\t" +
                        timePickerWakeUp.getMinute(), Toast.LENGTH_LONG).show();

                set_wake_up(timePickerWakeUp.getHour(), timePickerWakeUp.getMinute(), 0);
            }
        });
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu)
    {
        MenuInflater menuInflater = getMenuInflater();
        menuInflater.inflate(R.menu.options_menu, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item)
    {
        switch (item.getItemId())
        {
            case R.id.list_installed_apps:
                startActivity(new Intent(MainActivity.this, ListInstalledAppsActivity.class));
                return true;
            case R.id.show_device_ip:
                ShowIpDialog ipDialog = new ShowIpDialog();
                ipDialog.show(getFragmentManager(), TAG);
                return true;
        }
        return true;
    }

    private void get_admin_permission()
    {
        Intent intent = new Intent(DevicePolicyManager.ACTION_ADD_DEVICE_ADMIN);
        intent.putExtra(DevicePolicyManager.EXTRA_DEVICE_ADMIN, compName);
        intent.putExtra(DevicePolicyManager.EXTRA_ADD_EXPLANATION, "Additional text explaining why this needs to be added.");
        startActivityForResult(intent, 1);
    }

    private void set_wake_up(int hourOfDay, int minute, int second)
    {
        AlarmManager alarmManager = (AlarmManager) getSystemService(ALARM_SERVICE);
        Calendar calendar = Calendar.getInstance();
        calendar.set(Calendar.HOUR_OF_DAY, hourOfDay);
        calendar.set(Calendar.MINUTE, minute);
        calendar.set(Calendar.SECOND, second);

        // Debug.
        Log.d(TAG, calendar.getTime().toString());

        Intent intentAlarm = new Intent(this, AlarmReceiver.class);
        PendingIntent pendingIntent = PendingIntent.getBroadcast(this, 0, intentAlarm, 0);
        alarmManager.set(AlarmManager.RTC_WAKEUP, calendar.getTimeInMillis(), pendingIntent);
    }
}
