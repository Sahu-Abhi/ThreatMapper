import cx from 'classnames';
import { HiViewGridAdd } from 'react-icons/hi';
import { Button, Card, Step, Stepper, Typography } from 'ui-components';

import { CopyToClipboardIcon } from '@/components/CopyToClipboardIcon';
import { usePageNavigation } from '@/utils/usePageNavigation';

export const DockerConnectorForm = () => {
  const { navigate } = usePageNavigation();

  const code = `
docker run -dit \\
--cpus=".2" \\
--name=deepfence-agent \\
--restart on-failure \\
--pid=host \\
--net=host \\
--privileged=true \\
-v /sys/kernel/debug:/sys/kernel/debug:rw \\
-v /var/log/fenced \\
-v /var/run/docker.sock:/var/run/docker.sock \\
-v /:/fenced/mnt/host/:ro \\
-e USER_DEFINED_TAGS="" \\
-e MGMT_CONSOLE_URL="${window.location.host ?? '---CONSOLE-IP---'}" \\
-e MGMT_CONSOLE_PORT="443" \\
-e DEEPFENCE_KEY="${localStorage.getItem('dfApiKey') ?? '---DEEPFENCE-API-KEY---'}" \\
deepfenceio/deepfence_agent_ce:latest`;

  return (
    <Stepper>
      <Step indicator={<HiViewGridAdd />} title="Connect Docker Container">
        <div className={`${Typography.size.sm} dark:text-gray-200`}>
          Connect to Docker Container. Find out more information by{' '}
          <a
            href={`https://docs.deepfence.io/threatstryker/docs/sensors/docker/`}
            target="_blank"
            rel="noreferrer"
            className="text-blue-600 dark:text-blue-500 mt-2"
          >
            reading our documentation
          </a>
          .
        </div>
      </Step>
      <Step indicator="1" title="Deploy">
        <div className={`${Typography.size.sm} dark:text-gray-400`}>
          <p className="mb-2.5">
            Copy the following commands and paste them into your shell.
          </p>
          <Card className="w-full relative ">
            <pre
              className={cx(
                'pl-4 pt-4',
                'h-fit',
                `${Typography.weight.normal} ${Typography.size.xs} `,
              )}
            >
              {code}
            </pre>
            <CopyToClipboardIcon text={code} />
          </Card>
        </div>
      </Step>
    </Stepper>
  );
};