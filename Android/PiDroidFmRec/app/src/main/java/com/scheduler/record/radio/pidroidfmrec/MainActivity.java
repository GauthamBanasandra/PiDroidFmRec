package com.scheduler.record.radio.pidroidfmrec;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.Button;
import android.widget.TimePicker;
import android.widget.Toast;

public class MainActivity extends AppCompatActivity
{

    @Override
    protected void onCreate(Bundle savedInstanceState)
    {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        final TimePicker timePickerWakeUp = (TimePicker) findViewById(R.id.timePicker_wake_up);
        Button buttonSet = (Button) findViewById(R.id.button_set);

        assert buttonSet != null;
        buttonSet.setOnClickListener(new View.OnClickListener()
        {
            @Override
            public void onClick(View v)
            {
                assert timePickerWakeUp != null;
                Toast.makeText(MainActivity.this, timePickerWakeUp.getHour() + "\t" +
                        timePickerWakeUp.getMinute(), Toast.LENGTH_LONG).show();

            }
        });
    }
}
