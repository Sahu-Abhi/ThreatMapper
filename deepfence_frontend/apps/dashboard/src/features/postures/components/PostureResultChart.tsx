import { EChartsOption, PieSeriesOption, SeriesOption } from 'echarts';

import { ReactECharts } from '@/components/ReactEcharts';
import { Mode } from '@/theme/ThemeContext';

const option: EChartsOption = {
  backgroundColor: 'transparent',
  tooltip: {
    trigger: 'item',
  },
  legend: {
    show: false,
  },

  dataset: {
    source: [],
  },
  series: [
    {
      name: 'Postures',
      type: 'pie',
      radius: ['70%', '90%'],
      bottom: '10%',
      avoidLabelOverlap: true,
      label: {
        show: false,
        position: 'center',
      },
    },
  ],
};

export const PostureResultChart = ({
  theme,
  data,
  eoption,
}: {
  theme: Mode;
  data: { [key: string]: number };
  eoption?: EChartsOption;
}) => {
  if (!data) {
    return null;
  }

  option.dataset = {
    source: Object.keys(data).map((key) => ({
      Posture: key,
      value: data[key],
    })),
  };

  if (eoption && option && option.series) {
    const seriesOption = option.series as PieSeriesOption[];
    const seriesEOption = eoption.series as PieSeriesOption[];
    const merge = { ...seriesOption[0], ...seriesEOption[0] };
    option.series = [merge];
  }

  return <ReactECharts theme={theme === 'dark' ? 'dark' : 'light'} option={option} />;
};
